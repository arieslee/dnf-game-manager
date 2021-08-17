import dayjs from 'dayjs';
const timeHelper = {
    /**
     * 当前时间戳
     * @return <int>        unix时间戳(秒)
     */
    now: function () {
        return dayjs().unix();
    },
    /**
     * 转为正常的时间格式 年-月-日 时:分:秒
     * @param utcDateTime
     * @returns {string}
     */
    utcToTime(utcDateTime) {
        let date = new Date(utcDateTime);
        return dayjs(date).format("YYYY-MM-DD HH:mm:ss");
    },
    /**
     * 不太严谨的判断是否为时间戳
     * @param time
     * @returns {boolean|boolean}
     */
    isTimestamp(time) {
        return (/^\d+$/.test(time) && time + "".length === 10);
    },
    /**
     * 判断是否为日期+时间的格式,如2020-02-22 13:04:06
     * @param str
     * @returns {boolean|boolean}
     */
    isDateTime(str) {
        let reg = /^(\d{1,4})(-|\/)(\d{1,2})\2(\d{1,2}) (\d{1,2}):(\d{1,2}):(\d{1,2})$/;
        let r = str.toString().match(reg);
        if (r == null)
            return false;
        let d = new Date(r[1], r[3] - 1, r[4], r[5], r[6], r[7]);
        return (d.getFullYear() == r[1] && (d.getMonth() + 1) == r[3] &&
            d.getDate() == r[4] && d.getHours() == r[5] && d.getMinutes() == r[6] && d.getSeconds() == r[7]);
    },
    compare(start, end) {
        if (!this.isDateTime(start)) {
            start = this.toTimestamp(start);
        }
        if (!this.isDateTime(end)) {
            end = this.toTimestamp(end);
        }
        if (end <= start) {
            return false;
        }
        return true;
    },
    /**
     * 日期 转换为 Unix时间戳
     * @return <int>        unix时间戳(秒)
     * @param string 2014-01-01 20:20:20 or 2018-03-29T06:08:57.008Z  日期格式
     */
    toTimestamp: function (string) {
        if (!string) {
            return 0;
        }
        string = string.toISOString();
        if (string.indexOf('Z') !== -1 && string.indexOf('T') !== -1) {
            return (new Date(string)).getTime() / 1000;
        }
        return dayjs(string).unix();
    },
    /**
     * 时间戳转换日期
     * @param timestamp
     * @param format
     * @param timeZone
     */
    toDate: function (timestamp, format, timeZone) {
        if (!format) format = "YYYY-MM-DD";
        let _date = this.jsDate(timestamp, timeZone);
        return this.format(format, _date);
    },
    /**
     * 时间戳转换日期+时间
     * @param unixTime
     * @param format
     * @param timeZone
     * @returns {*}
     */
    toDateTime: function (unixTime, format, timeZone) {
        if (!format) format = "YYYY-MM-DD HH:mm:ss";
        let _date = this.jsDate(unixTime, timeZone);
        return this.format(format, _date);
    },
    /**
     * 根据unixTime获取js日期
     * @param unixTime
     * @param timeZone
     * @returns {Date}
     */
    jsDate: function (unixTime, timeZone) {
        if (typeof (timeZone) == 'number') {
            unixTime = parseInt(unixTime) + parseInt(timeZone) * 3600;
        }
        if (typeof unixTime !== 'string') {
            unixTime = '' + unixTime;
        }
        let len = unixTime.length;
        unixTime = (len <= 10) ? unixTime * 1000 : unixTime;
        return new Date(unixTime);
    },
    /**
     * 格式化时间
     * @param fmt
     * @param _date
     * @returns string
     */
    format: function (fmt, _date) {
        return dayjs(_date).format(fmt);
    },
    /**
     * 输出日期+时间
     * @param value
     * @returns {string|*}
     */
    dateTimeDisplay(value) {
        if (value > 0) {
            return this.toDateTime(value);
        }
        return '';
    },
    /**
     * 输出日期
     * @param value
     * @returns {string|*}
     */
    dateDisplay(value) {
        if (value > 0) {
            return this.toDate(value);
        }
        return '';
    },
}

export default timeHelper
