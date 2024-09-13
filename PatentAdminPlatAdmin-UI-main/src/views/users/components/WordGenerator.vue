<template>
  <div />
</template>
<script>
import Docxtemplater from 'docxtemplater'
import PizZip from 'pizzip'
import PizZipUtils from 'pizzip/utils/index'
import { saveAs } from 'file-saver'

function loadFile(url, callback) {
  PizZipUtils.getBinaryContent(url, callback)
}

export default {
  props: {
    dataList: {
      type: Object,
      default: () => ({})
    }
  },
  methods: {
    replaceErrors(key, value) {
      if (value instanceof Error) {
        return Object.getOwnPropertyNames(value).reduce(function(error, key) {
          error[key] = value[key]
          return error
        }, {})
      }
      return value
    },
    errorHandler(error) {
      console.log(JSON.stringify({ error: error }, this.replaceErrors()))
      if (error.properties && error.properties.errors instanceof Array) {
        const errorMessages = error.properties.errors
          .map(function(error) {
            return error.properties.explanation
          })
          .join('\n')
        console.log('errorMessages', errorMessages)
        // errorMessages is a humanly readable message looking like this :
        // 'The tag beginning with "foobar" is unopened'
      }
      throw error
    },
    renderDoc() {
      // 解决异步参数传递问题
      const _this = this
      loadFile('http://10.112.138.178:8000/static/docs/novelty_report_template.docx', function(
        error,
        content
      ) {
        if (error) {
          throw error
        }
        const zip = new PizZip(content)
        const doc = new Docxtemplater().loadZip(zip)
        doc.setData(_this.dataList)
        try {
          // render the document (replace all occurences of {first_name} by John, {last_name} by Doe, ...)
          doc.render()
        } catch (error) {
          // The error thrown here contains additional information when logged with JSON.stringify (it contains a properties object containing all suberrors).
          this.errorHandler(error)
        }
        const out = doc.getZip().generate({
          type: 'blob',
          mimeType:
            'application/vnd.openxmlformats-officedocument.wordprocessingml.document'
        }) // Output the document using Data-URI
        saveAs(out, 'output.docx')
      })
    }
  }
}
</script>
