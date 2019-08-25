import { CarplateModel } from '../models/carplateModel';

export class CarplatesApi {
    constructor(private readonly _baseUrl: string) {}

    getAll(): Promise<CarplateModel[]> {
        return new Promise((resolve, reject) => {
            fetch(`${this._baseUrl}/api/carplates`)
                .then(res => {
                    if (res.ok) {
                        return res.json();
                    }
                    reject(res);
                })
                .then(json => {
                    resolve(json);
                })
                .catch(err => {
                    reject(err);
                });
        });
    }

    get(id: string): Promise<CarplateModel> {
        return new Promise((resolve, reject) => {
            fetch(`${this._baseUrl}/api/carplates/${id}`)
                .then(res => {
                    if (res.ok) {
                        return res.json();
                    }
                    reject(res);
                })
                .then(json => {
                    resolve(json);
                })
                .catch(err => {
                    reject(err);
                });
        });
    }

    add(carplateModel: CarplateModel): Promise<void> {
        return new Promise((resolve, reject) => {
            fetch(
                new Request(`${this._baseUrl}/api/carplates`, {
                    method: 'post',
                    body: JSON.stringify(carplateModel)
                })
            )
                .then(res => {
                    if (res.ok) {
                        return res.json();
                    }
                    reject(res);
                })
                .then(json => {
                    resolve(json);
                })
                .catch(err => {
                    reject(err);
                });
        });
    }

    update(carplateModel: CarplateModel): Promise<void> {
        return new Promise((resolve, reject) => {
            fetch(
                new Request(`${this._baseUrl}/api/carplates`, {
                    method: 'put',
                    body: JSON.stringify(carplateModel)
                })
            )
                .then(res => {
                    if (res.ok) {
                        return res.json();
                    }
                    reject(res);
                })
                .then(json => {
                    resolve(json);
                })
                .catch(err => {
                    reject(err);
                });
        });
    }

    delete(id: string): Promise<void> {
        return new Promise((resolve, reject) => {
            fetch(
                new Request(`${this._baseUrl}/api/carplates/${id}`, {
                    method: 'delete'
                })
            )
                .then(res => {
                    if (res.ok) {
                        return res.json();
                    }
                    reject(res);
                })
                .then(json => {
                    resolve(json);
                })
                .catch(err => {
                    reject(err);
                });
        });
    }
}
