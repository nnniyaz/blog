import { currentTimeRFC3339 } from '@lib/time/time'
import { NextRequest } from 'next/server'

enum LogLevel {
    INFO = 'INFO',
    ERROR = 'ERROR',
    WARN = 'WARN',
}

class Logger {
    private log(req: NextRequest, res: Response, level: LogLevel) {
        console.log(
            `[${currentTimeRFC3339()}] [${req.headers.get('x-trace-id')}] [${req.headers.get('x-forwarded-for') || req.headers.get('x-real-ip')}] ${level.toUpperCase()} - ${req.method} ${req.nextUrl.pathname} ${res.status} - ${Date.now() - parseInt(req.headers.get('x-start-time') || '') || ''}ms`,
        )
    }

    info(req: NextRequest, res: Response) {
        this.log(req, res, LogLevel.INFO)
    }

    error(req: NextRequest, res: Response) {
        this.log(req, res, LogLevel.ERROR)
    }

    warn(req: NextRequest, res: Response) {
        this.log(req, res, LogLevel.WARN)
    }
}

let logger: Logger | null = null

export function getLogger() {
    if (!logger) {
        logger = new Logger()
    }

    return logger
}
