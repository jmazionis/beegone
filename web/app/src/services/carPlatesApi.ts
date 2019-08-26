import { CarplateModel } from '../models/carplateModel';

export interface CarplateValidationErrors {
    errors: string[]
}

export interface AddCarTemplateResponse {
    id: string;
}

export class CarplatesApi {
    constructor(private readonly _baseUrl: string) {}

    async getAll(): Promise<CarplateModel[]> {
        const response =  await fetch(`${this._baseUrl}/api/carplates`)
       
        if (!response.ok) {
            throw Error(response.statusText)            
        }                
        
        return response.json() 
    }

    async get(id: string): Promise<CarplateModel> {
        const response = await fetch(`${this._baseUrl}/api/carplates/${id}`)

        if (!response.ok) {
            throw Error(response.statusText)            
        }                
        
        return response.json() 
    }

    async add(carplateModel: CarplateModel): Promise<AddCarTemplateResponse> {
        const response =  await fetch(
            new Request(`${this._baseUrl}/api/carplates`, {
                method: 'post',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(carplateModel)
            })
        )
        
        if (!response.ok) {
            throw Error(response.statusText)            
        }  
        
        return response.json() as Promise<AddCarTemplateResponse>        
    }

    async update(carplateModel: CarplateModel): Promise<void> {
        const response = await fetch(
            new Request(`${this._baseUrl}/api/carplates`, {
                method: 'put',
                body: JSON.stringify(carplateModel)
            })
        )

        if (!response.ok) {
            throw Error(response.statusText)            
        }           
    }

   async delete(id: string): Promise<void> {
        const response = await fetch(
            new Request(`${this._baseUrl}/api/carplates/${id}`, {
                method: 'delete'
            })
        )    

        if (!response.ok) {
            throw Error(response.statusText)            
        }          
    }
}
