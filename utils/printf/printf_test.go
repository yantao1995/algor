package printf

import (
	"fmt"
	"testing"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func TestSill(t *testing.T) {
	// 创建新的表格实例
	tb := table.NewWriter()
	// 设置表头
	tb.AppendHeader(table.Row{"Date", "Today P&L", "Total P&L", "Today FEE", "Total FEE"})

	// 示例数据
	data := []struct {
		Date           int64
		TodayPnl       string
		PnlTday        string
		SourceFeeToday string
		SourceFeeTotal string
	}{
		{1637358400, "978.97", "-23714.45", "1205.42", "44912.76"},
		{1637262000, "730.55", "-24693.42", "3217.57", "43707.34"},
		{1637175600, "-940.63", "-25423.97", "3930.88", "40489.77"},
	}

	// 添加数据行
	for _, v := range data {
		tb.AppendRow(table.Row{
			time.Unix(v.Date, 0).Format("2006-01-02"),
			v.TodayPnl,
			v.PnlTday,
			v.SourceFeeToday,
			v.SourceFeeTotal,
		}, table.RowConfig{
			AutoMerge:      true,
			AutoMergeAlign: text.AlignCenter,
		})
	}
	tb.SetColumnConfigs([]table.ColumnConfig{{
		Number:   1,
		WidthMax: 12,
	}, {
		Number:   2,
		WidthMax: 12,
	}, {
		Number:   3,
		WidthMax: 12,
	}, {
		Number:   4,
		WidthMax: 12,
	}, {
		Number:   5,
		WidthMax: 12,
	}, {
		Number:   6,
		WidthMax: 12,
	}, {
		Number:   7,
		WidthMax: 12,
	}})
	// 渲染并输出表格
	fmt.Println(tb.Render() + time.Now().Format("\r\nquery_time : 2006-01-02 15:04:05"))
}
