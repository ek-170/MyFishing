export type Diary = {
    place?: string,
    caughtFish?: string[],
    comment?: string;
    rod?: string[],
    method?: string[],
    lure?: string[],
    weather?: string,
    wind?: number,
    tide?: string,
}

export type DiaryId = keyof Diary 