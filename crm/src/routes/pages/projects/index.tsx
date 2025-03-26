import { Table, TablePaginationConfig } from 'antd'
import { Card } from '@components/ui/card'
import { LinkButton } from '@components/ui/button/link-button.tsx'
import { RoutesPaths } from '@domain/base/routes-paths.ts'
import { translate } from '@lib/utils/translate.ts'
import { useTypedSelector } from '@lib/hooks/useTypedSelector.ts'
import { useActions } from '@lib/hooks/useActions.tsx'
import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { EmptyData } from '@components/ui/empty-data/empty-data.tsx'

export const Projects = () => {
    const { lang } = useTypedSelector((state) => state.system)
    const { projects, count } = useTypedSelector((state) => state.project)
    const [pagination, setPagination] = useState<TablePaginationConfig>({
        current: 1,
        pageSize: 10,
        total: count,
    })
    const {getProjects} = useActions()

    const columns = [
        {
            title: translate(lang, 'actions'),
            dataIndex: 'actions',
            key: 'actions',
        },
        {
            title: translate(lang, 'name'),
            dataIndex: 'name',
            key: 'name',
        },
    ]

    const data = projects.map((project) => ({
        key: project.id,
        actions: (
            <Link to={`${RoutesPaths.PROJECT_EDIT.replace(':id', project.id)}`}>
                {translate(lang, 'edit')}
            </Link>
        ),
        name: translate(lang, project.name),
    }))

    useEffect(() => {
        getProjects({
            offset: ((pagination.current || 1) - 1) * (pagination.pageSize || 10),
            limit: pagination.pageSize || 10,
        })

        // eslint-disable-next-line
    }, [pagination.current, pagination.pageSize])

    return (
        <div
            style={{
                display: 'flex',
                flexDirection: 'column',
                gap: '10px',
            }}
        >
            <Card
                title={'Add Your Projects'}
                subtitle={'Add your projects here to showcase them to the world.'}
            >
                <LinkButton href={RoutesPaths.PROJECT_ADD}>
                    {translate(lang, 'project_add')}
                </LinkButton>
            </Card>
            <Table
                bordered={true}
                columns={columns}
                dataSource={data}
                pagination={{
                    ...pagination,
                    onChange: (page) => {
                        setPagination({
                            ...pagination,
                            current: page,
                        })
                    },
                    onShowSizeChange: (_, size) => {
                        setPagination({
                            ...pagination,
                            pageSize: size,
                        })
                    },
                }}
                locale={{
                    emptyText: <EmptyData/>
                }}
            />
        </div>
    )
}
