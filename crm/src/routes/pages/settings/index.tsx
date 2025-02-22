import classes from './settings.module.scss'
import { Select } from 'antd'
import { Lang } from '@domain/base/ml-string.ts'
import { useTypedSelector } from '@lib/hooks/useTypedSelector.ts'
import { translate } from '@lib/utils/translate.ts'
import { useActions } from '@lib/hooks/useActions.tsx'

export const Settings = () => {
    const { lang, theme } = useTypedSelector((state) => state.system)
    const { changeLang, setTheme } = useActions()
    return (
        <>
            <div className={classes.settings__group}>
                <h3>{translate(lang, 'select_language')}</h3>
                <Select
                    onChange={(value) => changeLang(value)}
                    className={classes.select}
                    popupClassName={classes.select__dropdown}
                    defaultValue={lang}
                    options={
                        Object.values(Lang).map((lang) => ({
                            label: translate(lang, lang),
                            value: lang
                        }))
                    }
                />
            </div>

            <div className={classes.settings__group}>
                <h3>{translate(lang, 'select_theme')}</h3>
                {
                    ['light', 'dark'].map((t) => (
                        <div key={t} className={classes.radio__input__group}>
                            <input
                                className={classes.radio__input}
                                type={'radio'}
                                name={'theme'}
                                id={t}
                                value={t}
                                checked={theme === t}
                                onChange={() => setTheme(t as 'light' | 'dark')}
                            />
                            <label htmlFor={t}>
                                {translate(lang, t + '_theme')}
                            </label>
                        </div>
                    ))
                }
            </div>
        </>
    )
}
