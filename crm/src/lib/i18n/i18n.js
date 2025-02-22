import { parse } from 'csv-parse'
import fs from 'fs'
import path from 'path'

function TxtGen() {
    const csvFilePath = path.resolve('.', './txtmaps.csv')

    const headers = ['key', 'KZ', 'RU', 'EN']

    const fileContent = fs.readFileSync(csvFilePath, { encoding: 'utf-8' })

    console.log('generate txt from csv')

    parse(
        fileContent,
        {
            delimiter: ',',
            columns: headers,
        },
        (error, result) => {
            if (error) {
                console.error(error)

                process.exit(1)
            }
            const txtMap = {}
            result.forEach((v) => {
                txtMap[v.key] = v
            })
            fs.writeFile(
                './src/lib/i18n/i18ngen.ts',
                'export const txts: {[key: string]: {[key: string]: string}} = ' +
                    JSON.stringify(txtMap),
                (err) => {
                    if (err != null) {
                        console.log('error write to txt.json', err)
                        process.exit(1)
                    }
                },
            )
        },
    )
}

TxtGen()
