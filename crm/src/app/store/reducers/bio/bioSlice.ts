import { createSlice } from '@reduxjs/toolkit'

import { Bio } from '@domain/bio/bio'

type State = {
    bio: Bio | null
    isLoadingGetBio: boolean
}

const initialState: State = {
    bio: null,
    isLoadingGetBio: false,
}

type SetBioAction = {
    type: string
    payload: Bio | null
}

type SetIsLoadingGetBioAction = {
    type: string
    payload: boolean
}

export const bioSlice = createSlice({
    name: 'bio',
    initialState: initialState,
    reducers: {
        setBio: (state, action: SetBioAction) => {
            state.bio = action.payload
        },
        setIsLoadingGetBio: (state, action: SetIsLoadingGetBioAction) => {
            state.isLoadingGetBio = action.payload
        },
    },
})

export default bioSlice.reducer
