import {createSlice} from '@reduxjs/toolkit'

export enum Lang {
    EN = 'en',
    KZ = 'kz',
    RU = 'ru',
}

export const systemSlice = createSlice({
    name: 'system',
    initialState: {
        lang: Lang.EN,
    },
    reducers: {
        changeLang: (state, action) => {
            state.lang = action.payload
        }
    }
})

export default systemSlice.reducer
