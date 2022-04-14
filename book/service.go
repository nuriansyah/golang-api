package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	book, err := s.repository.FindAll()
	return book, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(BookRequest BookRequest) (Book, error) {
	price, _ := BookRequest.Price.Int64()
	rating, _ := BookRequest.Rating.Int64()
	discount, _ := BookRequest.Discount.Int64()
	book := Book{
		Title:       BookRequest.Title,
		Description: BookRequest.Description,
		Price:       int(price),
		Rating:      int(rating),
		Discount:    int(discount),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err

}

func (s *service) Update(ID int, BookRequest BookRequest) (Book, error) {
	book, err := s.repository.FindByID(ID)
	if err != nil {
		return book, err
	}

	price, _ := BookRequest.Price.Int64()
	rating, _ := BookRequest.Rating.Int64()
	discount, _ := BookRequest.Discount.Int64()

	book.Title = BookRequest.Title
	book.Description = BookRequest.Description
	book.Price = int(price)
	book.Rating = int(rating)
	book.Discount = int(discount)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	if err != nil {
		return book, err
	}
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
