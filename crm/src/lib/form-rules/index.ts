import { FormInstance, FormRule } from 'antd'
import { Lang, MlString } from '@domain/base/ml-string.ts'
import { txts } from '@lib/i18n/i18ngen.ts'

interface IRules {
    required: (message: string) => FormRule
    requiredI18n: (message: string) => FormRule
    email: (message: string) => FormRule
    minmaxLen: (message: string, min: number, max: number) => FormRule
    matchPass: (form: FormInstance, lang: Lang) => FormRule
    invalidAuth: (error: string | null, message: string) => FormRule
}

export const formRules: IRules = {
    required: (message: string) => ({ required: true, message: message }),
    requiredI18n: (message: string) => ({
        validator(_: unknown, value: MlString) {
            let countEmpties = 0
            Object.values(value).forEach((val) => {
                if (!val) {
                    countEmpties++
                }
            })
            if (countEmpties === Object.values(value).length) {
                return Promise.reject(message)
            }
            return Promise.resolve()
        },
    }),
    email: (message: string) => ({ type: 'email', message: message }),
    minmaxLen: (message: string, min: number, max: number) => ({
        validator(_: unknown, value: string) {
            if (!value) {
                return Promise.resolve()
            } else if (value.length < min || value.length > max) {
                return Promise.reject(message)
            }
            return Promise.resolve()
        },
    }),
    matchPass: (form, lang) => ({
        validator(_, value) {
            if (!value || form.getFieldValue('password') === value) {
                return Promise.resolve()
            }
            return Promise.reject(new Error(txts.password_does_not_match[lang]))
        },
    }),
    invalidAuth: (error: string | null, message: string) => ({
        validator() {
            if (error) {
                return Promise.reject(message)
            }
            return Promise.resolve()
        },
    }),
}
