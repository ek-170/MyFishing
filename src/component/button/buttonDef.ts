export type ButtonType = 'submit' | 'reset' | 'button' | undefined

export type ButtonMode = 'positive' | 'negative'

export type ButtonProps<T extends ButtonType, P extends ButtonMode> = {
    type: T,
    mode:P,
    children:React.ReactNode,
    disabled?: boolean,
    hidden?: boolean,
    onClick?: any,
}