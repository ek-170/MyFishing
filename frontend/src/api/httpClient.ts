import { Diary } from 'model/diary'
import { DIARYURL } from 'utils/const/endpoint'
import axios from 'axios';

export const fetchDiary = async (id: string) => {
    try {
        const res = await axios.get<Diary>(DIARYURL + id)
        return res.data
    } catch (err) {
        if(axios.isAxiosError(err)){
            console.error(err.message)
            return err
        }
    }
};

export const createDiary = async (d: Diary) => {
    try {
        const res = await axios.post<never>(DIARYURL, d)
        return res.statusText
    } catch (err) {
        if(axios.isAxiosError(err)){
            console.error(err.message)
            return err
        }
    }
};

export const updateDiary =async (id: string, d: Diary) => {
    try {
        const res = await axios.put<Diary>(DIARYURL + id + '/edit', d)
        return res.data
    } catch (err) {
        if(axios.isAxiosError(err)){
            console.error(err.message)
            return err
        }
    }
};

export const deleteDiary = async (id: string) => {
    try {
        const res = await axios.delete<never>(DIARYURL + id)
        return res.statusText
    } catch (err) {
        if(axios.isAxiosError(err)){
            console.error(err.message)
            return err
        }
    }
};