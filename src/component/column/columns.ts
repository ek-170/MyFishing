import { TextColumn, NumberColumn, PicklistColumn, MultiPicklistColumn, TextareaColumn } from 'component/column/columnDef';


export const textColumns:TextColumn[] = [
    {
        label: '場所',
        id: 'place',
        type: 'input',
        input: {
            type: 'text',
            name: 'place',
        },
        value: undefined,
    }
]

export const numberColumns:NumberColumn[] = [
    {
        label: '風速',
        id: 'wind',
        type: 'input',
        input: {
            type: 'number',
            name: 'wind',
        },
        value: undefined,
    }
]

export const picklistColumns: PicklistColumn[] = [
    {
        label: '天気',
        id: 'weather',
        type: 'picklist',
        select:{
            name: 'weather',
        },
        options: [
            {
                value: 'sunny',
                text: '晴れ',
            },
            {
                value: 'cloudy',
                text: '曇り',
            },
            {
                value: 'rainy',
                text: '雨',
            },
        ],
        value: undefined,
    },
    {
        label: '潮',
        id: 'tide',
        type: 'picklist',
        select:{
            name: 'tide',
        },
        options: [
            {
                value: 'ooshio',
                text: '大潮',
            },
            {
                value: 'nakashio',
                text: '中潮',
            },
            {
                value: 'koshio',
                text: '小潮',
            },
            {
                value: 'wakashio',
                text: '若潮',
            },
            {
                value: 'nagashio',
                text: '長潮',
            },
        ],
        value: undefined,
    },
]

export const multiPicklistColumns: MultiPicklistColumn[] = [
    {
        label: '釣れた魚',
        id: 'caughtFish',
        type: 'multipicklist',
        select:{
            name: 'caughtFish',
            required: true,
        },
        options: [{
                value: 'aji',
                text: 'アジ',
            },
            {
                value: 'buri',
                text: 'ブリ',
            },
            {
                value: 'kasago',
                text: 'カサゴ',
            },
            {
                value: 'tachiuo',
                text: 'タチウオ',
            },
            {
                value: 'haze',
                text: 'ハゼ',
            }
        ],
        value: undefined,
    },
    {
        label: '使用ロッド',
        id: 'rod',
        type: 'multipicklist',
        select:{
            name: 'rod',
            required: true,
        },
        options: [{
                value: 'soare_extune_MB',
                text: 'ソアレ エクスチェーン MB',
            },
            {
                value: 'nessa_BB',
                text: 'ネッサ BB',
            },
            {
                value: 'hardrocker_SS',
                text: 'ハードロッカー SS',
            },
            {
                value: 'burenias',
                text: 'ブレニアス',
            },
            {
                value: 'luminous',
                text: 'ルミナス',
            }
        ],
        value: undefined,
    },
    {
        label: '釣法',
        id: 'method',
        type: 'multipicklist',
        select:{
            name: 'method',
            required: true,
        },
        options: [{
                value: 'bukkomi',
                text: 'ぶっこみ釣り',
            },
            {
                value: 'sabiki',
                text: 'サビキ釣り',
            },
            {
                value: 'lure',
                text: 'ルアー釣り',
            },
            {
                value: 'ana',
                text: '穴釣り',
            },
            {
                value: 'uki',
                text: 'ウキ釣り',
            }
        ],
        value: undefined,
    },
    {
        label: 'ルアー',
        id: 'lure',
        type: 'multipicklist',
        select:{
            name: 'lure',
            required: true,
        },
        options: [{
                value: 'CLA016',
                text: 'CLA016',
            },
            {
                value: 'CLA017',
                text: 'CLA017',
            },
            {
                value: 'CLB022',
                text: 'CLB022',
            },
            {
                value: 'silentassasin',
                text: 'サイレントアサシン',
            },
            {
                value: 'gurubin',
                text: 'グルービン',
            }
        ],
        value: undefined,
    },
]

export const textareaColumns: TextareaColumn[] = [
{
    label: 'コメント',
    id: 'comment',
    type: 'textarea',
    textarea: {
        name: 'comment',
    },
    value: undefined,
}
]