import { createSlice } from '@reduxjs/toolkit'

import { Lang } from '@domain/base/ml-string'

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

type SetIsLoadingChangeLangAction = {
    type: string
    payload: boolean
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
        setIsLoadingChangeLang: (
            state,
            action: SetIsLoadingChangeLangAction,
        ) => {
            state.isLoadingChangeLang = action.payload
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
