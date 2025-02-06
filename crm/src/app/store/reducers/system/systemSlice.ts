import { createSlice } from '@reduxjs/toolkit'

import { Lang } from '@domain/base/mlString/mlString'

type State = {
    lang: Lang
    isLoadingChangeLang: boolean
}

const initialState: State = {
    lang: Lang.EN,
    isLoadingChangeLang: false,
}

type ChangeLangAction = {
    type: string
    payload: Lang
}

type SetIsLoadingChangeLangAction = {
    type: string
    payload: boolean
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
    },
})

export default systemSlice.reducer
