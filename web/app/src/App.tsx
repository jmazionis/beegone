import React from 'react';
import './App.css';
import { CarplatesApi } from './services/carPlatesApi';
import { CarplateTable } from './components/CarplateTable';

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
            <CarplateTable carPlatesApi={props.carplatesApi} />
        </div>
    );
};

export default App;
