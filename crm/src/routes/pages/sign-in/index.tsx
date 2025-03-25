import classes from './sign-in.module.scss'
import { Form } from 'antd'
import { translate } from '@lib/utils/translate.ts'
import { useTypedSelector } from '@lib/hooks/useTypedSelector.ts'
import { Button } from '@components/ui/button'
import { Input } from '@components/ui/input'
import { EyeInvisibleOutlined, EyeOutlined } from '@ant-design/icons'
import { useState } from 'react'
import { formRules } from '@lib/form-rules'
import { useActions } from '@lib/hooks/useActions.tsx'

export const SignIn = () => {
    const { lang } = useTypedSelector((state) => state.system)
    const { isLoadingSignIn } = useTypedSelector((state) => state.auth)
    const { signIn } = useActions()
    const [showPassword, setShowPassword] = useState(false)

    const onFinish = (values: { email: string; password: string }) => {
        signIn(values)
    }

    return (
        <main className={classes.sign_in}>
            <div className={classes.sign_in__block}>
                <h1 className={classes.sign_in__heading}>
                    {translate(lang, 'blog')}
                </h1>
                <Form layout={'vertical'} noValidate={true} onFinish={onFinish}>
                    <Form.Item<string>
                        name={'email'}
                        label={translate(lang, 'email')}
                        required={false}
                        rules={[
                            formRules.required(
                                translate(lang, 'please_type_email'),
                            ),
                            formRules.email(
                                translate(lang, 'type_valid_email'),
                            ),
                        ]}
                    >
                        <Input
                            type={'email'}
                            placeholder={translate(lang, 'type_your_email')}
                        />
                    </Form.Item>
                    <Form.Item<string>
                        name={'password'}
                        label={translate(lang, 'password')}
                        required={false}
                        rules={[
                            formRules.required(
                                translate(lang, 'please_type_your_password'),
                            ),
                        ]}
                    >
                        <Input
                            type={showPassword ? 'text' : 'password'}
                            placeholder={translate(lang, 'type_your_password')}
                            icon={{
                                element: showPassword ? (
                                    <EyeInvisibleOutlined />
                                ) : (
                                    <EyeOutlined />
                                ),
                                onClick: () => setShowPassword(!showPassword),
                            }}
                        />
                    </Form.Item>
                    <Form.Item>
                        <Button type={'submit'} loading={isLoadingSignIn}>
                            {translate(lang, 'sign_in')}
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        </main>
    )
}
