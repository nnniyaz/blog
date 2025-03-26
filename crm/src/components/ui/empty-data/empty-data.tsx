import { InboxOutlined } from '@ant-design/icons'
import { translate } from '@lib/utils/translate.ts'
import { useTypedSelector } from '@lib/hooks/useTypedSelector.ts'

export const EmptyData = () => {
    const { lang } = useTypedSelector((state) => state.system)
    return (
        <div
            style={{
                height: '200px',
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'center',
                gap: '10px',
            }}
        >
            <InboxOutlined style={{ fontSize: '60px' }} />
            <p style={{ fontSize: '16px', fontWeight: '500', margin: '0' }}>
                {translate(lang, 'no_data')}
            </p>
        </div>
    )
}
