import React from 'react';
import './App.css';
import { CarplatesApi } from './services/carPlatesApi';
import { CarplatesTable } from './components/CarplatesTable';

interface AppProps {
    carplatesApi: CarplatesApi;
}

const App: React.FunctionComponent<AppProps> = (props: AppProps) => {
    return (
        <div
            style={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                marginTop: '100px'
            }}
        >
            <CarplatesTable carPlatesApi={props.carplatesApi} />
        </div>
    );
};

export default App;
