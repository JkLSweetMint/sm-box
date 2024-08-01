package url_shortner_srv

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/url_shortner/objects/constructors"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/tracer"
	pb "sm-box/transport/proto/pb/golang/url_shortner-service"
)

// Create - создание сокращенного url.
func (srv *server) Create(ctx context.Context, request *pb.UrlShortnerCreateRequest) (response *pb.UrlShortnerCreateResponse, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelTransportGrpc)

		trc.FunctionCall(ctx, request)
		defer func() { trc.Error(err).FunctionCallFinished(response) }()
	}

	response = new(pb.UrlShortnerCreateResponse)

	// Проверка данных
	{
		if request.Properties == nil {
			srv.components.Logger.Error().
				Text("Invalid arguments were received. ").Write()

			var cErr = common_errors.InvalidArguments()
			cErr.Details().Set("properties", "Is empty. ")

			err = cErr
			return
		}
	}

	var url *models.ShortUrlInfo

	// Обработка
	{
		var constructor *constructors.ShortUrl

		// Подготовка конструктора
		{
			constructor = &constructors.ShortUrl{
				Source: request.Source,

				Accesses: &constructors.ShortUrlAccesses{
					RolesID:       make([]common_types.ID, 0),
					PermissionsID: make([]common_types.ID, 0),
				},
				Properties: &constructors.ShortUrlProperties{
					Type:         types.ShortUrlType(request.Properties.Type),
					NumberOfUses: request.Properties.NumberOfUses,
					StartActive:  request.Properties.StartActive.AsTime(),
					EndActive:    request.Properties.EndActive.AsTime(),
					Active:       request.Properties.Active,
				},
			}
			constructor.FillEmptyFields()

			if request.Accesses != nil {
				for _, id := range request.Accesses.RolesID {
					constructor.Accesses.RolesID = append(constructor.Accesses.RolesID, common_types.ID(id))
				}

				for _, id := range request.Accesses.PermissionsID {
					constructor.Accesses.PermissionsID = append(constructor.Accesses.PermissionsID, common_types.ID(id))
				}
			}
		}

		if url, err = srv.controllers.UrlsManagement.Create(ctx, constructor); err != nil {
			srv.components.Logger.Error().
				Format("Could not create short url: '%s'. ", err).Write()

			return
		}
	}

	// Преобразование данных в структуры grpc
	{
		response.UrlShortner = &pb.UrlShortner{
			ID:        uint64(url.ID),
			Source:    url.Source,
			Reduction: url.Reduction,

			Accesses: &pb.UrlShortner_XAccesses{
				RolesID:       make([]uint64, 0),
				PermissionsID: make([]uint64, 0),
			},
			Properties: &pb.UrlShortner_XProperties{
				Type:   string(url.Properties.Type),
				Active: url.Properties.Active,

				NumberOfUses:         url.Properties.NumberOfUses,
				RemainedNumberOfUses: url.Properties.RemainedNumberOfUses,

				StartActive: timestamppb.New(url.Properties.StartActive),
				EndActive:   timestamppb.New(url.Properties.EndActive),
			},
		}

		if url.Accesses != nil {
			for _, id := range url.Accesses.RolesID {
				response.UrlShortner.Accesses.RolesID = append(response.UrlShortner.Accesses.RolesID, uint64(id))
			}
			for _, id := range url.Accesses.PermissionsID {
				response.UrlShortner.Accesses.PermissionsID = append(response.UrlShortner.Accesses.PermissionsID, uint64(id))
			}
		}
	}

	return
}
