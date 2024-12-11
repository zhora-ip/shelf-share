package store

type Store interface {
	User() UserRepository
	Book() BookRepository
	Library() LibraryRepository
	Feedback() FeedbackRepository
	Discussion() DiscussionRepository
}
