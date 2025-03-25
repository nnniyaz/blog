import { createSlice } from '@reduxjs/toolkit'

import { Bio } from '@domain/bio/bio.ts'

type State = {
    bio: Bio | null
}

const initialState: State = {
    bio: null,
}

type SetBioAction = {
    type: string
    payload: Bio | null
}

export const bioSlice = createSlice({
    name: 'bio',
    initialState: initialState,
    reducers: {
        setBio: (state, action: SetBioAction) => {
            state.bio = action.payload
        },
    },
})

export default bioSlice.reducer
