import { Options, MultiOptions } from 'component/column/input/pickListValueDef'
import { DiaryId } from 'dataType/diary'

export type ValueType = string | string[] | number
export type BaseColumnProps = {
    label: string,
    id: ColumnId,
    value?: ValueType,
}

type HandleChange = {
    handleChange?: any
}

export type Picklist = {
    select: {
        name: string,
        required?: boolean,
    },
    options: Options[],
}

export type MultiPicklist = {
    select: {
        name: string,
        required?: boolean,
    },
    options: MultiOptions[],
}

export type InputType = 'text' | 'number'
export type Input<T extends InputType> = {
    input: {
        type:T,
        name: string,
        required?: boolean,
        placeholder?: string,
        disabled?: boolean,
        maxlength?: number,
    },
    value?: T extends 'text'
        ? string
        : T extends 'number'
        ? number
        : never,
}

export type Textarea = {
    textarea: {
        name: string,
        required?: boolean,
        placeolder?: string,
        rows?: number,
        cols?: number,
        disabled?: boolean,
        maxlength?: number,
    },
}


export type PicklistProps = Picklist & BaseColumnProps & HandleChange
export type MultiPicklistProps = MultiPicklist & BaseColumnProps & HandleChange
export type InputProps<T extends InputType> = Input<T> & BaseColumnProps & HandleChange
export type TextareaProps = Textarea & BaseColumnProps & HandleChange

type ColumnType = { type: 'input' | 'picklist' | 'multipicklist' |'textarea' }
export type Column = ( TextColumn | NumberColumn | PicklistColumn | MultiPicklistColumn | TextareaColumn );
export type TextColumn = InputProps<'text'> & ColumnType
export type NumberColumn = InputProps<'number'> & ColumnType
export type PicklistColumn = PicklistProps & ColumnType
export type MultiPicklistColumn = MultiPicklistProps & ColumnType
export type TextareaColumn = TextareaProps & ColumnType

type ColumnId = DiaryId