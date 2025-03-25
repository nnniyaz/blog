import { createSlice } from '@reduxjs/toolkit'

import { Lang } from '@domain/base/ml-string.ts'

type State = {
    lang: Lang
    isLoadingChangeLang: boolean
    theme: 'light' | 'dark'
}

const initialState: State = {
    lang: Lang.EN,
    isLoadingChangeLang: false,
    theme: 'light',
}

type ChangeLangAction = {
    type: string
    payload: Lang
}

type SetThemeAction = {
    type: string
    payload: 'light' | 'dark'
}

export const systemSlice = createSlice({
    name: 'system',
    initialState: initialState,
    reducers: {
        changeLang: (state, action: ChangeLangAction) => {
            state.lang = action.payload
        },
        setTheme: (state, action: SetThemeAction) => {
            if (action.payload === 'light') {
                document.body.classList.remove('dark-theme')
            } else {
                document.body.classList.add('dark-theme')
            }
            state.theme = action.payload
        },
    },
})

export default systemSlice.reducer
