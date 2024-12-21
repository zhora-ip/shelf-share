BEGIN;
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_trigger WHERE tgname='feedback_after_insert') THEN
        DROP TRIGGER feedback_after_insert ON feedback;
    END IF;
END $$;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_proc WHERE proname='update_average_grade') THEN 
        DROP FUNCTION update_average_grade();
    END IF;
END $$;


DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='books' AND column_name='grade_sum') THEN
        ALTER TABLE books DROP COLUMN grade_sum;
    END IF;

    IF EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='books' AND column_name='grade_count') THEN
        ALTER TABLE books DROP COLUMN grade_count;
    END IF;
END $$;

COMMIT;
