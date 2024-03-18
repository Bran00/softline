"use client"
import { useState, useEffect } from 'react';
import Link from 'next/link';
import axios from 'axios';
import styles from './Cliente.module.css';

const Cliente = () => {
  const [clientes, setClientes] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    const fetchClientes = async () => {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/clientes', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        setClientes(response.data);
      } catch (error) {
        console.error('Erro ao obter lista de clientes:', error);
      }
    };

    fetchClientes();
  }, []);

  const handleSearch = (e) => {
    setSearchTerm(e.target.value);
  };

  const filteredClientes = clientes.filter((cliente) =>
    cliente.nome.toLowerCase().includes(searchTerm.toLowerCase()) ||
    cliente.documento.toLowerCase().includes(searchTerm.toLowerCase()) || cliente.fantasia.toLowerCase().includes(searchTerm.toLowerCase) || cliente.documento.includes(searchTerm) || cliente.endereco.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className={styles.container}>
      <Link className={styles.button} href="/">Home</Link>
      <h1>Lista de Clientes</h1>
      <Link className={styles.button} href="/clienteCadastro">Cadastro de Clientes</Link>
      <input
        type="text"
        placeholder="Buscar cliente..."
        value={searchTerm}
        onChange={handleSearch}
      />
      <ul>
        {filteredClientes.map((cliente, index) => (
        <Link href={`/listaClientes/${cliente.id}`}>
          <li key={cliente.id} className={index % 2 === 0 ? styles.zebra : ''}>
              
                <strong>Nome:</strong> {cliente.nome}<br />
                <strong>Fantasia:</strong> {cliente.fantasia}<br />
                <strong>Documento:</strong> {cliente.documento}<br />
                <strong>Endereço:</strong> {cliente.endereco}
              
          </li>
        </Link>
        ))}
      </ul>
    </div>
  );
};

export default Cliente;
