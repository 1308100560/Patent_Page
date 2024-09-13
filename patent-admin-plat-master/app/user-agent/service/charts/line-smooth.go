package charts

const lineSmoothProfile = `{
  "xAxis": {
    "type": "category",
    "data": $CATE,
	$ROTATE
  },
  "yAxis": {
    "type": "value"
  },
  "series": [
    {
      "data": $DATA,
      "type": "line",
      "smooth": true
    }
  ],
  "toolbox": {
    "feature": {
      "dataView": {
        "show": "true",
        "title": "数据视图",
        "lang": ["数据视图", "关闭", "刷新"]
      },
      "saveAsImage": {
        "show": true,
        "title": "导出"
      }
    }
  }
}`

func genLineSmoothProfile(cate []string, data []int, isRotate bool) string {
	p := newProfile(lineSmoothProfile)
	if isRotate {
		p = p.replace("$ROTATE", ROTATE)
	} else {
		p = p.replace("$ROTATE", "").
			replace("$CATE,", "$CATE")
	}
	cateTemp := strListTemplate(cate)
	dataTemp := intListTemplate(data)

	return p.replace("$CATE", cateTemp).
		replace("$DATA", dataTemp).
		String()
}
