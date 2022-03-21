import { Diary } from 'dataType/diary'
import { DateForInput } from 'component/column/columnDef'

export const testData: Diary = {
    date: '2022-04-01' as DateForInput,
    place: '大洗漁港',
    caughtFish: ['aji','kasago'],
    comment: 'テストコメントです。',
    rod:['soare_extune_MB'],
    method:['bukkomi','ana'],
    lure:['CLA017','silentassasin'],
    weather: 'rainy',
    wind: 5,
    tide: 'koshio',
}