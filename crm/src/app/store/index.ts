import {configureStore} from '@reduxjs/toolkit'

import systemReducer from './reducers/system/systemSlice'

export const store = configureStore({
    reducer: {
        system: systemReducer,
    }
})

export type RootState = ReturnType<typeof store.getState>
