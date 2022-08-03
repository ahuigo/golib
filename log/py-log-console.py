import logging
import datetime as dt
import sys
import json

#logging.basicConfig()
logging.root.setLevel(logging.INFO)
logging.basicConfig(level=logging.INFO)


class MLogFormatter(logging.Formatter):
    converter=dt.datetime.fromtimestamp
    def formatTime(self, record, datefmt=None):
        ct = self.converter(record.created)
        if datefmt:
            s = ct.astimezone().strftime(datefmt)
        else:
            s = ct.astimezone().strftime("%Y-%m-%d %H:%M:%S.%f%z")
            #s = "%s.%03d" % (t, record.msecs)
        return s


def getLogger(name='root', level=logging.INFO):
    formatmsg='%(asctime)s\t%(levelname)s\t%(name)s\t%(pathname)s:%(lineno)s\t%(message)s'
    datefmt='%Y-%m-%dT%I:%M:%S.%f%z'
    #logging.basicConfig(format=formatmsg, datefmt=datefmt, level=level)
    #logging.basicConfig(format=formatter, datefmt=datefmt, level=level)
    # formatter
    formatter = MLogFormatter(fmt=formatmsg, datefmt=datefmt)
    for handler in logging.root.handlers[:]:
        logging.root.removeHandler(handler)

    ch = logging.StreamHandler()
    #ch = logging.StreamHandler(sys.stdout)
    ch.setLevel(level)
    #ch.setFormatter(logging.Formatter(formatmsg))
    ch.setFormatter(formatter)

    logger = logging.getLogger(name)
    logger.handlers.clear()
    logger.addHandler(ch)
    logger.propagate = False
    return logger


logger = getLogger('translog') # 实例名
log = {
    # 日志类型 1.数据格式错误　DATAERROR 2.网络服务错误    NETWORKERROR 3.计算错误             COMPUTEERROR
    'type': "DATAERROR", 
    # 日志消息
    'msg':'消息:\t\nM1903-SKU4-V03-2006-0064_2020_08_27_13_29_11 　starts preparing file',
    # 其它：自定义
    "pod_name":'detection-partition-64c469f6c8-q9zlb',
    'module_name':'translog-A[73]',
}

logger.error("ERROR1")
logger.error("[MONITOR_LOG_TYPE=DATAERROR]  数据错误...")
logger.info("logger.info...")
#------------------------------
#2020-11-05T08:11:51+0800	INFO	translog	py-log-console.py:30	{"type": "DATAERROR", "msg": "M1903-SKU4-V03-2006-0064_2020_08_27_13_29_11 \u3000starts preparing file", "pod_name": "detection-partition-64c469f6c8-q9zlb", "module_name": "translog-A[73]"}
