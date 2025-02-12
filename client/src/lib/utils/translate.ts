import { Lang, MlString } from '@domain/base/ml-string'
import { txts } from '@lib/i18n/i18ngen'

export function translate(lang: Lang, key: string | MlString): string {
    if (typeof key === 'string') {
        const translation = txts?.[key]?.[lang.toUpperCase()] || null
        if (!translation) {
            return 'Translation not found'
        }
        return translation
    } else {
        return key[lang] || 'Translation not found'
    }
}
