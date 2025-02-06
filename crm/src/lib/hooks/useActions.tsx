import { useDispatch } from 'react-redux'
import { bindActionCreators } from 'redux'

import { systemSlice } from '@app/store/reducers/system/systemSlice'

export const useActions = () => {
    const dispatch = useDispatch()
    return bindActionCreators(
        {
            ...systemSlice.actions,
        },
        dispatch,
    )
}
