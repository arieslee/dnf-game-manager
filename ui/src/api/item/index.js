import {get} from '@/utils/request';

export const getItemList = (params) => {
    return get('item',{params:params});
}