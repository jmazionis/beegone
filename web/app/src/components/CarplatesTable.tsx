import React, { Component } from 'react';
import { CarplatesApi } from '../services/carPlatesApi';
import MaterialTable from 'material-table';
import { CarplateModel } from '../models/carplateModel';

export interface ICarplatesTableProps {
    carPlatesApi: CarplatesApi;
}

export interface ICarplatesTableState {
    isLoading: boolean;
    carplates: CarplateModel[];
}

export class CarplatesTable extends Component<
    ICarplatesTableProps,
    ICarplatesTableState
> {
    constructor(props: ICarplatesTableProps) {
        super(props);

        this.state = { isLoading: false, carplates: [] };
    }

    async componentWillMount() {
        try {
            this.setState({
                isLoading: true
            });
            const carplates = await this.props.carPlatesApi.getAll();
            this.setState({
                carplates
            });
        } catch (error) {
            console.log('Error fetching carplates: ', error);
        } finally {
            this.setState({
                isLoading: false
            });
        }
    }

    render() {
        const { isLoading, carplates } = this.state;
        return (
            <MaterialTable
                style={{
                    minWidth: '70%'
                }}
                isLoading={isLoading}
                columns={[
                    { title: 'Plate Id', field: 'plateId' },
                    { title: 'Model', field: 'modelName' },
                    { title: 'Year', field: 'modelYear' },
                    { title: 'Owner', field: 'owner' }
                ]}
                data={carplates}
                title="Carplates App"
            />
        );
    }
}
