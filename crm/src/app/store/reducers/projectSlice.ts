import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'

import { Project } from '@domain/project/project.ts'
import { ProjectService } from '@http/project/projectService.ts'
import { RootState } from '@app/store'
import { PaginationParams } from '@domain/base/pagination-params.ts'

type State = {
    projects: Project[]
    count: number
    isLoadingGetProjects: boolean
    selectedProject: Project | null
    isLoadingGetSelectedProject: boolean
}

const initialState: State = {
    projects: [],
    count: 0,
    isLoadingGetProjects: false,
    selectedProject: null,
    isLoadingGetSelectedProject: false,
}

type SetProjectsAction = {
    type: string
    payload: {projects: Project[] | null, count: number}
}

type SetSelectedProjectAction = {
    type: string
    payload: Project | null
}

export const getProjects = createAsyncThunk(
    'project/getProjects',
    async (params: PaginationParams, {getState, dispatch}) => {
        const state = getState() as RootState
        try {
            const res = await ProjectService.getAll({
                lang: state.system.lang,
                body: params,
            })
            if (res.data.success) {
                dispatch(projectSlice.actions.setProjects(res.data.data))
            }
        } catch (error) {
            throw error
        }
    }
)

export const projectSlice = createSlice({
    name: 'project',
    initialState: initialState,
    reducers: {
        setProjects: (state, action: SetProjectsAction) => {
            state.projects = action.payload.projects || []
            state.count = action.payload.count
        },
        setSelectedProject: (state, action: SetSelectedProjectAction) => {
            state.selectedProject = action.payload
        },
    },
})

export default projectSlice.reducer
