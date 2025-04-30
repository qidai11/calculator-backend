// frontend/__tests__/Calculator.test.tsx
import { render, screen, fireEvent, waitFor } from '@testing-library/react'
import Calculator from '../pages/index'
import { CalculatorService } from '../gen/calculator/v1/calculator_connectweb'
import { OperationRequest } from '../gen/calculator/v1/calculator_pb'

// Mock Connect 服务
jest.mock('../gen/calculator/v1/calculator_connectweb', () => ({
    CalculatorService: {
        add: jest.fn().mockResolvedValue({ result: 5 }),
        subtract: jest.fn().mockResolvedValue({ result: 2 }),
        multiply: jest.fn().mockResolvedValue({ result: 10 }),
        divide: jest.fn().mockImplementation((req: OperationRequest) => {
            if (req.b === 0) throw new Error('division by zero');
            return { result: req.a / req.b };
        })
    }
}));

describe('Calculator Component', () => {
    beforeEach(() => {
        jest.clearAllMocks()
    })

    test('应正确渲染计算器表单元素', () => {
        render(<Calculator />)
        expect(screen.getAllByRole('spinbutton')).toHaveLength(2)
        expect(screen.getByRole('combobox')).toBeInTheDocument()
        expect(screen.getByText('计算结果')).toBeInTheDocument()
    })

    test('应正确处理加法运算请求', async () => {
        render(<Calculator />)
        
        fireEvent.change(screen.getAllByRole('spinbutton')[0], { target: { value: '2' } })
        fireEvent.change(screen.getAllByRole('spinbutton')[1], { target: { value: '3' } })
        fireEvent.click(screen.getByText('计算结果'))

        await waitFor(() => {
            expect(CalculatorService.add).toHaveBeenCalledWith(
                new OperationRequest({ a: 2, b: 3 })
            )
            expect(screen.getByText('结果: 5.00')).toBeInTheDocument()
    })
    })

    test('应捕获除零异常并显示错误信息', async () => {
        (CalculatorService.divide as jest.Mock).mockRejectedValueOnce(
        new Error('division by zero')
        )

        render(<Calculator />)
        
        fireEvent.change(screen.getAllByRole('spinbutton')[0], { target: { value: '10' } })
        fireEvent.change(screen.getAllByRole('spinbutton')[1], { target: { value: '0' } })
        fireEvent.change(screen.getByRole('combobox'), { target: { value: '/' } })
        fireEvent.click(screen.getByText('计算结果'))

        await waitFor(() => {
        expect(screen.getByText(/计算失败/)).toBeInTheDocument()
        })
    })

    test('应验证数字输入有效性', async () => {
        render(<Calculator />)
        
        // Test invalid input
        fireEvent.change(screen.getAllByRole('spinbutton')[0], { target: { value: 'abc' } })
        fireEvent.click(screen.getByText('计算结果'))

        expect(await screen.findByText('请输入有效数字')).toBeInTheDocument()
    })

    test('应正确切换运算符类型', async () => {
        render(<Calculator />)
        
        fireEvent.change(screen.getByRole('combobox'), { target: { value: '*' } })
        fireEvent.change(screen.getAllByRole('spinbutton')[0], { target: { value: '5' } })
        fireEvent.change(screen.getAllByRole('spinbutton')[1], { target: { value: '2' } })
        fireEvent.click(screen.getByText('计算结果'))

        await waitFor(() => {
        expect(CalculatorService.multiply).toHaveBeenCalledWith(
            new OperationRequest({ a: 5, b: 2 })
        )
        })
    })
})