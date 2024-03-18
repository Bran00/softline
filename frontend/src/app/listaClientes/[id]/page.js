"use client"
import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import axios from 'axios';
import Link from 'next/link';
import styles from './ClienteDetalhes.module.css'

const ClienteDetalhes = ({ params }) => {
  const router = useRouter()
  const [cliente, setCliente] = useState(null);
  const [error, setError] = useState(null);
  const [nome, setNome] = useState('');
  const [fantasia, setFantasia] = useState('');
  const [documento, setDocumento] = useState('');
  const [endereco, setEndereco] = useState('');
  const [showConfirmation, setShowConfirmation] = useState(false);

  useEffect(() => {
    const fetchCliente = async () => {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get(`http://localhost:8080/clientes/${params.id}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setCliente(response.data);
        setNome(response.data.nome);
        setFantasia(response.data.fantasia);
        setDocumento(response.data.documento);
        setEndereco(response.data.endereco);
      } catch (error) {
        setError(error);
      }
    };

    fetchCliente();
  }, [params.id]);

  const handleUpdateCliente = async () => {
    const clienteData = {
      nome,
      fantasia,
      documento,
      endereco,
    };

    try {
      const token = localStorage.getItem('token');
      const response = await axios.put(`http://localhost:8080/clientes/${params.id}`, clienteData, {
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });
      alert('Cliente atualizado');
      router.replace("/listaClientes")
    } catch (error) {
      console.error('Erro ao atualizar cliente:', error);
    }
  };

  const handleDeleteConfirmation = () => {
    setShowConfirmation(true);
  };

  const handleDeleteCliente = async () => {
    try {
      const token = localStorage.getItem('token');
      await axios.delete(`http://localhost:8080/clientes/${params.id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      alert('Cliente excluído com sucesso!');
      router.replace("/listaClientes")
    } catch (error) {
      console.error('Erro ao excluir cliente:', error);
    }
  };

  if (error) {
    return <div>Erro ao obter detalhes do cliente: {error.message}</div>;
  }

  if (!cliente) {
    return <div>Carregando...</div>;
  }

  return (
    <div className={styles.container}>
      <Link href="/">Home</Link>
      <h1>Detalhes do Cliente</h1>
      <label>
        Nome:
        <input type="text" value={nome} onChange={(e) => setNome(e.target.value)} />
      </label>
      <label>
        Fantasia:
        <input type="text" value={fantasia} onChange={(e) => setFantasia(e.target.value)} />
      </label>
      <label>
        Documento:
        <input type="text" value={documento} onChange={(e) => setDocumento(e.target.value)} />
      </label>
      <label>
        Endereço:
        <input type="text" value={endereco} onChange={(e) => setEndereco(e.target.value)} />
      </label>
      <button onClick={handleUpdateCliente} className={styles.update}>Atualizar Cliente</button>
      <button onClick={handleDeleteConfirmation}className={styles.delete}>Excluir Cliente</button>
      {showConfirmation && (
        <div className={styles.confirmation}>
          <p>Deseja realmente excluir o cliente {cliente.nome}?</p>
          <button onClick={handleDeleteCliente}className={styles.red}>Sim</button>
          <button onClick={() => setShowConfirmation(false)}
            className={styles.green}
          >Cancelar</button>
        </div>
      )}
    </div>
  );
};

export default ClienteDetalhes;
