package container

import (
	"github.com/miyantara7/desent-solutions/internal/handler"
	"github.com/miyantara7/desent-solutions/internal/router"
	repo "github.com/miyantara7/desent-solutions/internal/shared/repository/store"
	"github.com/miyantara7/desent-solutions/internal/shared/repository/usecase"
	bookStore "github.com/miyantara7/desent-solutions/internal/store/book"
	bookUsecase "github.com/miyantara7/desent-solutions/internal/usecase/book"
	"github.com/sarulabs/di/v2"
)

const (
	ContainerRepository = "book-repository"
	ContainerUsecase    = "book-usecase"
	ContainerHandler    = "book-handler"
	ContainerRouter     = "router"
)

func Build() (ctn di.Container, err error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return
	}

	builder.Add(di.Def{
		Name: ContainerRepository,
		Build: func(ctn di.Container) (interface{}, error) {
			return bookStore.NewBookRepository(), nil
		},
	})

	builder.Add(di.Def{
		Name: ContainerUsecase,
		Build: func(ctn di.Container) (interface{}, error) {
			repo := ctn.Get(ContainerRepository).(repo.BookRepository)
			return bookUsecase.NewBookUsecase(repo), nil
		},
	})

	builder.Add(di.Def{
		Name: ContainerHandler,
		Build: func(ctn di.Container) (interface{}, error) {
			uc := ctn.Get(ContainerUsecase).(usecase.BookUsecase)
			return handler.NewBookHandler(uc), nil
		},
	})

	builder.Add(di.Def{
		Name: ContainerRouter,
		Build: func(ctn di.Container) (interface{}, error) {

			bookHandler := ctn.Get(ContainerHandler).(*handler.BookHandler)

			bookRoute := router.NewBookRoute(bookHandler)
			authRoute := router.NewAuthRoute(bookHandler)
			healthRoute := router.NewHealthRoute()

			return router.SetupRouter(
				bookRoute,
				authRoute,
				healthRoute,
			), nil
		},
	})

	return builder.Build(), nil
}
