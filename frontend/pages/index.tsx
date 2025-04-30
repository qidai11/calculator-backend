import { useState } from 'react';
import { createPromiseClient } from '@bufbuild/connect';
import { createConnectTransport } from '@bufbuild/connect-web';
import { CalculatorService } from '../gen/calculator/v1/calculator_connectweb';
import { OperationRequest } from '../gen/calculator/v1/calculator_pb';
import styles from '../styles/Calculator.module.css'; // 模块化样式导入

const transport = createConnectTransport({
  baseUrl: 'http://localhost:8080',
});

const client = createPromiseClient(CalculatorService, transport);

export default function Calculator() {
  const [a, setA] = useState('');
  const [b, setB] = useState('');
  const [operator, setOperator] = useState('+');
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState('');

  const handleCalculate = async () => {
    setError('');
    setResult(null);

    const numA = parseFloat(a);
    const numB = parseFloat(b);
    if (isNaN(numA) || isNaN(numB)) {
      setError('请输入有效数字');
      return;
    }

    try {
      const request = new OperationRequest({ a: numA, b: numB });
      let response;
      
      switch (operator) {
        case '+': response = await client.add(request); break;
        case '-': response = await client.subtract(request); break;
        case '*': response = await client.multiply(request); break;
        case '/': response = await client.divide(request); break;
        default: setError('未知运算符'); return;
      }

      setResult(response.result);
    } catch (err) {
      setError(`计算失败: ${err instanceof Error ? err.message : '未知错误'}`);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.inputGroup}>
        <input
          type="number"
          value={a}
          onChange={(e) => setA(e.target.value)}
          placeholder="输入数字"
          step="any"
          className={styles.numberInput}
        />
        
        <select
          value={operator}
          onChange={(e) => setOperator(e.target.value)}
          className={styles.operatorSelect}
        >
          <option value="+">+</option>
          <option value="-">-</option>
          <option value="*">×</option>
          <option value="/">÷</option>
        </select>

        <input
          type="number"
          value={b}
          onChange={(e) => setB(e.target.value)}
          placeholder="输入数字"
          step="any"
          className={styles.numberInput}
        />
      </div>

      <button 
        onClick={handleCalculate} 
        className={styles.calculateButton}
      >
        计算结果
      </button>

      {result !== null && (
        <div className={styles.resultBox}>
          结果: {result.toFixed(2)}
        </div>
      )}

      {error && <div className={styles.errorBox}>{error}</div>}
    </div>
  );
}