import dayjs from 'dayjs';
import 'dayjs/locale/ja';

export enum DateFormat {
  YYYY_MM_DD = 'YYYY-MM-DD',
}

export const formatDate = (date: Date, type: DateFormat): string => {
  switch (type) {
    case DateFormat.YYYY_MM_DD:
      return dayjs(date)
        .locale('ja')
        .format('YYYY-MM-DD'); 
    default:
      return '';
  }
};