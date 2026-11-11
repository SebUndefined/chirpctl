package cmdutils

import (
	"github.com/sebundefined/chirpctl/internal/service/shared"
	"github.com/spf13/cobra"
)

func AttachPaginationFlags(cmd *cobra.Command, p *shared.ListPaginationRequest) {
	cmd.Flags().Uint32Var(&p.Limit, "limit", 100, "Page size")
	cmd.Flags().Uint32Var(&p.Offset, "offset", 0, "Offset for pagination")
	cmd.Flags().StringVar(&p.Search, "search", "", "Search query")
}

func AttachOrderByFlags(cmd *cobra.Command, p *shared.ListOrderByRequest) {
	cmd.Flags().StringVar(&p.OrderBy, "order-by", "", "Order by")
	cmd.Flags().BoolVar(&p.OrderByDesc, "order-by-desc", true, "Order by descending")
}

func AttachTagFlags(cmd *cobra.Command, p *shared.ListTagsFilterRequest) {
	cmd.Flags().StringToStringVar(
		&p.Tags,
		"tag",
		nil,
		"Filter by tags (key=value[,key=value...])",
	)
}
