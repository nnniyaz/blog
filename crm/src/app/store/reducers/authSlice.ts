import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { AuthService } from '@http/auth/authService.ts'
import { RootState } from '@app/store'
import {
    currentUserSlice,
    getCurrentUser,
} from '@app/store/reducers/currentUserSlice.ts'

type State = {
    isAuth: boolean
    isLoadingSignIn: boolean
    isLoadingSignOut: boolean
    error: string | null
}

const initialState: State = {
    isAuth: false,
    isLoadingSignIn: false,
    isLoadingSignOut: false,
    error: null,
}

type SetIsAuthAction = {
    type: string
    payload: boolean
}

interface SignInIn {
    email: string
    password: string
}

export const signIn = createAsyncThunk(
    'auth/signIn',
    async (data: SignInIn, { getState, dispatch }) => {
        const state = getState() as RootState
        try {
            const res = await AuthService.signIn({
                lang: state.system.lang,
                body: data,
            })
            if (res.data.success) {
                dispatch(getCurrentUser())
            }
        } catch (error) {
            throw error
        }
    },
)

export const signOut = createAsyncThunk(
    'auth/signOut',
    async (_, { getState, dispatch }) => {
        const state = getState() as RootState
        try {
            const res = await AuthService.signOut({
                lang: state.system.lang,
                body: null,
            })
            if (res.data.success) {
                dispatch(currentUserSlice.actions.setCurrentUser(null))
            }
        } catch (error) {
            throw error
        }
    },
)

export const authSlice = createSlice({
    name: 'auth',
    initialState: initialState,
    reducers: {
        setIsAuth: (state, action: SetIsAuthAction) => {
            state.isAuth = action.payload
        },
    },
    extraReducers: (builder) => {
        builder.addCase(signIn.pending, (state) => {
            state.isLoadingSignIn = true
            state.error = null
        })
        builder.addCase(signIn.fulfilled, (state) => {
            state.isLoadingSignIn = false
        })
        builder.addCase(signIn.rejected, (state, action) => {
            state.isLoadingSignIn = false
            state.error = action.payload as string
        })
        builder.addCase(signOut.pending, (state) => {
            state.isLoadingSignOut = true
            state.error = null
        })
        builder.addCase(signOut.fulfilled, (state) => {
            state.isLoadingSignOut = false
            state.isAuth = false
        })
        builder.addCase(signOut.rejected, (state, action) => {
            state.isLoadingSignOut = false
            state.error = action.payload as string
        })
    },
})

export default authSlice.reducer
