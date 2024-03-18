"use client"
import { useState } from 'react';
import Link from 'next/link';
import styles from './CadastroCliente.module.css';

const ClienteCadastro = () => {
  const [nome, setNome] = useState('');
  const [fantasia, setFantasia] = useState('');
  const [documento, setDocumento] = useState('');
  const [endereco, setEndereco] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    const token = localStorage.getItem('token');

    const clienteData = {
      nome,
      fantasia,
      documento,
      endereco
    };

    try {
      const response = await fetch('http://localhost:8080/clientes', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(clienteData)
      });

      if (response.ok) {
        setNome('');
        setFantasia('');
        setDocumento('');
        setEndereco('');

        alert('Cliente cadastrado com sucesso!');
      } else {
        alert('Falha ao cadastrar cliente. Por favor, tente novamente.');
      }
    } catch (error) {
      console.error('Erro ao cadastrar cliente:', error);
    }
  };

  return (
    <div className={styles.container}>
      <h1>Cadastro de Cliente</h1>
      <form onSubmit={handleSubmit}>
        <input type="text" placeholder="Nome" value={nome} onChange={(e) => setNome(e.target.value)} />
        <input type="text" placeholder="Fantasia" value={fantasia} onChange={(e) => setFantasia(e.target.value)} />
        <input type="text" placeholder="Documento" value={documento} onChange={(e) => setDocumento(e.target.value)} />
        <input type="text" placeholder="Endereço" value={endereco} onChange={(e) => setEndereco(e.target.value)} />
        <button type="submit">Cadastrar</button>
      </form>
      <Link href="/listaClientes">Lista de Clientes</Link>
    </div>
  );
};

export default ClienteCadastro;
