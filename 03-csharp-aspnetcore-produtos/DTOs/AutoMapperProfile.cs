using AutoMapper;
using ProdutosApi.Models;

namespace ProdutosApi.DTOs
{
    public class AutoMapperProfile : Profile
    {
        public AutoMapperProfile()
        {
            CreateMap<Produto, ProdutoResponseDTO>();
            CreateMap<ProdutoRequestDTO, Produto>()
                .ForMember(dest => dest.Id, opt => opt.Ignore())
                .ForMember(dest => dest.DataCriacao, opt => opt.MapFrom(src => DateTime.UtcNow));
        }
    }
}