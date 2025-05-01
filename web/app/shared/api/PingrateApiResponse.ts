export class PingrateApiResponse<T = undefined> {
    public constructor(
        private _ok: boolean,
        private _data?: T,
        private _error?: Error,
    ) {}

    public get ok(): boolean {
        return this._ok;
    }

    public get data(): T | undefined {
        return this._data;
    }

    public get error(): Error | undefined {
        return this._error;
    }
}