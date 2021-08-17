import {get,post} from '@/utils/request';

/**
 * 登录
 */
export const login = (params) => {
  let url = 'login';
  return post(url, params);
}
/**
 * 请求id
 */
export const requestId = () => {
  let url = 'request-id';
  return get(url);
}