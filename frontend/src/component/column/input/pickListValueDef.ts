export type Options = Weather | Tide
type Weather = {
    value:'none'
        |'sunny'
        | 'cloudy'
        | 'rainy'
    text:''
        | '晴れ'
        | '曇り'
        | '雨',
    selected?: boolean,
}

type Tide = {
    value:'none'
        | 'ooshio'
        | 'nakashio'
        | 'koshio'
        | 'wakashio'
        | 'nagashio'
    text:''
        | '大潮'
        | '中潮'
        | '小潮'
        | '若潮'
        | '長潮'
    selected?: boolean,
}


export type MultiOptions = Fish | Rod | Lure | Method
type Fish = {
    value:'aji'
        | 'buri'
        | 'kasago'
        | 'tachiuo'
        | 'haze' ,
    text: 'アジ'
        | 'ブリ'
        | 'カサゴ'
        | 'タチウオ'
        | 'ハゼ' ,
    selected?: boolean,
}

type Rod = {
    value:'soare_extune_MB'
        | 'nessa_BB'
        | 'hardrocker_SS'
        | 'burenias'
        | 'luminous' ,
    text: 'ソアレ エクスチェーン MB'
        | 'ネッサ BB'
        | 'ハードロッカー SS'
        | 'ブレニアス'
        | 'ルミナス' ,
    selected?: boolean,
}

type Method = {
    value:'bukkomi'
        | 'sabiki'
        | 'lure'
        | 'ana'
        | 'uki' ,
    text: 'ぶっこみ釣り'
        | 'サビキ釣り'
        | 'ルアー釣り'
        | '穴釣り'
        | 'ウキ釣り' ,
    selected?: boolean,
}

type Lure = {
    value:'CLA016'
        | 'CLA017'
        | 'CLB022'
        | 'silentassasin'
        | 'gurubin',
    text: 'CLA016'
        | 'CLA017'
        | 'CLB022'
        | 'サイレントアサシン'
        | 'グルービン' ,
    selected?: boolean,
}