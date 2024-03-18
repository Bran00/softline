"use client"
import { useEffect, useState } from 'react';
import Link from 'next/link';
import axios from 'axios';
import styles from './Produto.module.css';
import { useRouter } from 'next/router';

const Produto = () => {
  const [produtos, setProdutos] = useState([]);
  const [error, setError] = useState(null);
  const [searchTerm, setSearchTerm] = useState('');

  useEffect(() => {
    const fetchProdutos = async () => {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/produtos', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        setProdutos(response.data);
      } catch (error) {
        setError(error);
      }
    };

    fetchProdutos();
  }, []);

  const handleSearch = (e) => {
    setSearchTerm(e.target.value);
  };

  const filteredProdutos = produtos.filter((produto) =>
    produto.nome.toLowerCase().includes(searchTerm.toLowerCase()) || produto.descricao.toLowerCase().includes(searchTerm.toLocaleLowerCase) || produto.codigo_barras.includes(searchTerm)
  );

  if (error) {
    return <div>Erro ao obter lista de produtos: {error.message}</div>;
  }

  return (
    <div className={styles.container}>
      <Link className={styles.button} href="/">Home</Link>
      <h1>Lista de Produtos</h1>
      <Link className={styles.button} href="/produtoCadastro">Cadastro de Produtos</Link>
      <input
        type="text"
        placeholder="Buscar produto..."
        value={searchTerm}
        onChange={handleSearch}
      />
      <ul className={styles.produtoList}>
        {filteredProdutos.map((produto) => (
          <Link href={`/listaProdutos/${produto.id}`}>
          <li key={produto.id} className={styles.zebra}>
           
              
                <strong>Nome:</strong> {produto.nome}<br />
                <strong>Descrição:</strong> {produto.descricao}<br />
                <strong>Código de Barras:</strong> {produto.codigo_barras}<br />
                <strong>Valor de Venda:</strong> R$ {produto.valor_venda}<br />
                <strong>Peso Bruto:</strong> {produto.peso_bruto}<br />
                <strong>Peso Líquido:</strong> {produto.peso_liquido}<br />
          </li>

            </Link>
          
        ))}
      </ul>
    </div>
  );
};

export default Produto;
