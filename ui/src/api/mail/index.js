import {get,post} from '@/utils/request';

export const sendMail = (params)=>{
    return post("/mail",params);
}