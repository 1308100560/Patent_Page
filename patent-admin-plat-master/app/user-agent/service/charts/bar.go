package charts

const barProfile = `{
  "tooltip": {
    "trigger": "axis",
    "axisPointer": {
      "type": "shadow"
    }
  },
  "grid": {
    "left": "3%",
    "right": "4%",
    "bottom": "3%",
    "containLabel": true
  },
  "xAxis": [
    {
      "type": "category",
      "data": $CATE,
      $ROTATE
    }
  ],
  "yAxis": [
    {
      "type": "value"
    }
  ],
  "series": [
    {
      "name": "Direct",
      "type": "bar",
      "barWidth": "60%",
      "data": $DATA
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

const ROTATE = `
      "axisTick": {
        "alignWithLabel": true
      },
      "axisLabel": {
        "interval": 0,
        "rotate": 45
      }`

func genBarProfile(cate []string, data []int, isRotate bool) string {
	p := newProfile(barProfile)
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
