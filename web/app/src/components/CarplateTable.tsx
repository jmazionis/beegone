import React, { Component } from 'react';
import { CarplatesApi } from '../services/carPlatesApi';
import MaterialTable from 'material-table';
import { CarplateModel } from '../models/carplateModel';

export interface ICarplateTableProps {
    carPlatesApi: CarplatesApi;
}

export interface ICarplateTableState {
    isLoading: boolean;
    carplates: CarplateModel[];
}

export class CarplateTable extends Component<
    ICarplateTableProps,
    ICarplateTableState
> {
    constructor(props: ICarplateTableProps) {
        super(props);

        this.state = { isLoading: false, carplates: [] };
    }

    async componentDidMount() {
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

    handleAddCarplate = async (carplate: CarplateModel): Promise<void> => {
        const { carPlatesApi } = this.props;
        try {
            this.setState({
                isLoading: true
            });

            const addCarTemplateResponse = await carPlatesApi.add(carplate);
            carplate.id = addCarTemplateResponse.id;
            
            this.setState(prevState => {
                return { carplates: [...prevState.carplates, carplate] };
            });            
        } catch (error) {
            console.log('Error adding carplate: ', error);
        } finally {
            this.setState({
                isLoading: false
            });
        }
    };

    handleUpdateCarplate = async (carplate: CarplateModel): Promise<void> => {
        const { carPlatesApi } = this.props;
        try {
            this.setState({
                isLoading: true
            });
            await carPlatesApi.update(carplate);
            this.setState(prevState => {
                return {
                    carplates: prevState.carplates.map(c => {
                        if (c.id === carplate.id) {
                            return carplate;
                        }
                        return c;
                    })
                };
            });
        } catch (error) {
            console.log('Error updating carplate: ', error);
        } finally {
            this.setState({
                isLoading: false
            });
        }
    };

    handleDeleteCarplate = async (carplate: CarplateModel): Promise<void> => {
        const { carPlatesApi } = this.props;
        try {
            this.setState({
                isLoading: true
            });
            await carPlatesApi.delete(carplate.id as string);
            this.setState(prevState => {
                return {
                    carplates: prevState.carplates.filter(c => {
                        return c.id !== carplate.id;
                    })
                };
            });
        } catch (error) {
            console.log('Error deleting carplate: ', error);
        } finally {
            this.setState({
                isLoading: false
            });
        }
    };

    render() {
        const { isLoading, carplates } = this.state;
        return (
            <MaterialTable
                style={{
                    minWidth: '70%'
                }}
                title="Carplates App"
                isLoading={isLoading}
                columns={[
                    { title: 'Plate Id', field: 'plateId' },
                    { title: 'Model', field: 'modelName' },
                    { title: 'Year', field: 'modelYear' },
                    { title: 'Owner', field: 'owner' }
                ]}
                data={carplates}
                editable={{
                    onRowAdd: carplate => this.handleAddCarplate(carplate),
                    onRowUpdate: carplate =>
                        this.handleUpdateCarplate(carplate),
                    onRowDelete: carplate => this.handleDeleteCarplate(carplate)
                }}
            />
        );
    }
}
