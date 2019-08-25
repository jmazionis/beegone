import React, { Component } from 'react';
import { CarplatesApi } from '../services/carPlatesApi';
import MaterialTable from 'material-table';

export interface ICarplatesTableProps {
    carPlatesApi: CarplatesApi;
}

export interface ICarplatesTableState {}

export class CarplatesTable extends Component<
    ICarplatesTableProps,
    ICarplatesTableState
> {
    async componentWillMount() {
        const carplates = await this.props.carPlatesApi.getAll();
        console.log(carplates);
    }

    constructor(props: ICarplatesTableProps) {
        super(props);

        this.state = {};
    }

    public render() {
        return (
            <MaterialTable
                style={{
                    minWidth: '70%'
                }}
                columns={[
                    { title: 'Plate Id', field: 'plateId' },
                    { title: 'Model', field: 'modelName' },
                    { title: 'Year', field: 'modelYear' },
                    { title: 'Owner', field: 'owner' }
                ]}
                data={[
                    {
                        plateId: 'Mehmet',
                        modelName: 'Baran',
                        modelYear: 1990,
                        owner: 'Driver'
                    }
                ]}
                title="Carplates App"
            />
        );
    }
}
