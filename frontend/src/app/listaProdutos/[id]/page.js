"use client"
import { useEffect, useState } from 'react';
import Link from 'next/link';
import axios from 'axios';
import styles from './ProdutoDetalhes.module.css';
import { useRouter } from 'next/navigation';

const ProdutoDetalhes = ({ params }) => {
  const router = useRouter()
  const [produto, setProduto] = useState(null);
  const [error, setError] = useState(null);
  const [showConfirmation, setShowConfirmation] = useState(false);

  useEffect(() => {
    const fetchProduto = async () => {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get(`http://localhost:8080/produtos/${params.id}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setProduto(response.data);
      } catch (error) {
        setError(error);
      }
    };

    fetchProduto();
  }, [params.id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setProduto(prevState => ({
      ...prevState,
      [name]: value
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const token = localStorage.getItem('token');
      await axios.put(`http://localhost:8080/produtos/${params.id}`, produto, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      alert('Produto atualizado com sucesso!');
      router.replace("/listaProdutos")
    } catch (error) {
      setError(error);
    }
  };

  const handleDeleteConfirmation = () => {
    setShowConfirmation(true);
  };

  const handleDeleteProduto = async () => {
    try {
      const token = localStorage.getItem('token');
      await axios.delete(`http://localhost:8080/produtos/${params.id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      alert('Produto excluído com sucesso!');
      router.replace("/listaProdutos")
    } catch (error) {
      setError(error);
    }
  };

  if (error) {
    return <div>Erro ao obter detalhes do produto: {error.message}</div>;
  }

  if (!produto) {
    return <div>Carregando...</div>;
  }

  return (
    <div className={styles.container}>
      <Link className={styles.button} href="/">Home</Link>
      <h1>Detalhes do Produto</h1>
      <form onSubmit={handleSubmit}>
        <label>Nome: </label> 
        <input type="text" name="nome" value={produto.nome} onChange={handleChange} />
        <label>Descrição: </label> 
        <input type="text" name="descricao" value={produto.descricao} onChange={handleChange} />
        <label>Código de Barras: </label> 
        <input type="text" name="codigo_barras" value={produto.codigo_barras} onChange={handleChange} />
        <label>Valor de Venda: </label> 
        <input type="number" name="valor_venda" value={produto.valor_venda} onChange={handleChange} />
        <label>Peso Bruto em KG: </label> 
        <input type="number" name="peso_bruto" value={produto.peso_bruto} onChange={handleChange} />
        <label>Peso Liquido em KG: </label> 
        <input type="number" name="peso_liquido" value={produto.peso_liquido} onChange={handleChange} />
        <button type="submit">Atualizar</button>
      </form>
      <button onClick={handleDeleteConfirmation}>Excluir</button>
      {showConfirmation && (
        <div className={styles.confirmation}>
          <p>Deseja realmente excluir o produto {produto.nome}?</p>
          <button onClick={handleDeleteProduto}>Sim</button>
          <button onClick={() => setShowConfirmation(false)}>Cancelar</button>
        </div>
      )}
    </div>
  );
};

export default ProdutoDetalhes;
