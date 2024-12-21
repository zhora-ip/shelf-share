DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='books' AND column_name='grade_sum') THEN
        ALTER TABLE books ADD COLUMN grade_sum FLOAT DEFAULT 0;
    END IF;

    IF NOT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='books' AND column_name='grade_count') THEN
        ALTER TABLE books ADD COLUMN grade_count INT DEFAULT 0;
    END IF;
END $$;


DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE table_name='feedback' AND constraint_name='check_grade_constraint') THEN
        ALTER TABLE feedback ADD CONSTRAINT check_grade_constraint CHECK (grade >= 1 AND grade <= 5);
    END IF;
END $$;


UPDATE books b
SET
    grade_sum = COALESCE(sub.sum_grade, 0),
    grade_count = COALESCE(sub.count_grade, 0)
FROM (
    SELECT
        book_id,
        SUM(grade) AS sum_grade,
        COUNT(*) AS count_grade
    FROM
        feedback
    GROUP BY
        book_id
) sub
WHERE b.id = sub.book_id;


CREATE OR REPLACE FUNCTION update_average_grade()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE books
    SET
        grade_sum = grade_sum + NEW.grade,
        grade_count = grade_count + 1,
        avg_grade = (grade_sum + NEW.grade) / (grade_count + 1)
    WHERE id = NEW.book_id;    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_trigger WHERE tgname='feedback_after_insert') THEN
        CREATE TRIGGER feedback_after_insert
        AFTER INSERT ON feedback
        FOR EACH ROW
        EXECUTE FUNCTION update_average_grade();
    END IF;
END $$;
