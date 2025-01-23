import {createSlice} from '@reduxjs/toolkit'
import {Lang} from "../../../../domain/base/mlString/mlString";

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
