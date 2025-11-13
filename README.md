# Khởi đầu của Overlix

Chào mừng bạn đến với repository onboarding nội bộ\!

Repository này giúp **nhân viên mới nhanh chóng hiểu kiến trúc hệ thống, chuẩn code, workflow Git và best practices** của công ty.

> ⚠️ Tài liệu này chưa đầy đủ và đang được bổ sung.

---

## Mục Lục

- [Bắt đầu nhanh](#bắt-đầu-nhanh)
- [Tổng quan kiến trúc](#tổng-quan-kiến-trúc)
- [Quy chuẩn Git & Commit](#quy-chuẩn-git--commit)
- [Quy trình phát triển (Workflow Git)](#quy-trình-phát-triển-workflow-git-)
- [Quy chuẩn Code & Best Practices](#quy-chuẩn-code--best-practices-)
- [Ví dụ / Tutorial](#ví-dụ--tutorial)

---

## Bắt đầu nhanh

1.  Clone repository:

    ```bash
    git clone https://git.overlix.net/overlix/onboarding.git
    cd onboarding
    ```

2.  Tham khảo các ví dụ trong thư mục `examples/`:

    - Rust: `examples/rust`
    - Golang: `examples/golang`
    - Flutter: `examples/flutter`
    - React: `examples/react`

3.  Đọc tài liệu module trong `ARCHITECTURE.md` để hiểu cấu trúc hệ thống.

---

## Sơ lược kiến trúc

Hệ thống bao gồm:

- **Application (Mobile/Desktop)**: Ứng dụng dành cho người dùng cuối, chịu trách nhiệm cung cấp giao diện và trải nghiệm đầy đủ.
- **Website (Web App)**: Trang web cho người dùng, phục vụ cho việc truy cập từ trình duyệt, cung cấp chức năng tương tự ứng dụng.
- **API (Backend)**: Ứng dụng để giao diện người dùng tương tác với phía máy chủ. Được xây dựng theo kiến trúc **Microservices** (sẽ được mô tả chi tiết trong `ARCHITECTURE.md`).

---

## Quy chuẩn Git & Commit

Commit message cần theo cấu trúc như sau (dựa trên **Conventional Commits**):

```text
<type>(optional scope): <description>

[optional body]

[optional footer]
```

| Type         | Ý nghĩa                                                                             | Phiên bản (Semantic Versioning) |
| :----------- | :---------------------------------------------------------------------------------- | :------------------------------ |
| **fix**      | Sửa lỗi.                                                                            | `PATCH`                         |
| **feat**     | Thêm tính năng mới.                                                                 | `MINOR`                         |
| **refactor** | Refactor code, không thêm tính năng hay sửa lỗi.                                    | Không ảnh hưởng                 |
| **chore**    | Thay đổi hỗ trợ, cấu hình, không ảnh hưởng chức năng (e.g., cập nhật dependencies). | Không ảnh hưởng                 |
| **test**     | Thêm hoặc sửa test.                                                                 | Không ảnh hưởng                 |
| **docs**     | Các thay đổi chỉ liên quan đến documentation.                                       | Không ảnh hưởng                 |
| **perf**     | Các thay đổi code nhằm tăng hiệu năng.                                              | Không ảnh hưởng                 |
| **ci**       | Thay đổi đối với configuration và scripts file CI.                                  | Không ảnh hưởng                 |

- **BREAKING CHANGE**: Một commit có chữ `BREAKING CHANGE:` ở phần đầu của `[optional body]` hoặc phần footer sẽ giới thiệu những **thay đổi API bị phá vỡ** (Tương đương với `MAJOR` trong semantic versioning). Một BREAKING CHANGE có thể là một phần commits của bất kì `<type>` nào.
- **Scope**: (Tùy chọn) Dùng để chỉ định phần nào của dự án bị ảnh hưởng (ví dụ: `(api)`, `(auth)`, `(ui-header)`).

### Ví dụ

#### Commit message với mô tả và breaking change trong body

```
feat: thêm tính năng subgroup vào api

Thêm endpoint mới để quản lý subgroup cho người dùng.

BREAKING CHANGE: tính năng `subgroup` giới thiệu nhiều endpoints mới, yêu cầu client cập nhật thư viện giao tiếp.
```

#### Commit message mà không cần body

```
docs: sửa chính tả cho CHANGELOG
```

#### Commit message với scope

```
feat(lang): thêm ngôn ngữ tiếng Nhật
```

---

## Quy trình phát triển (Workflow Git)

Chúng ta sử dụng một biến thể của **Git Flow** đơn giản hóa, tập trung vào nhánh chính và nhánh tính năng/sửa lỗi.

1.  **Nhánh Chính (Primary Branch)**:

    - Sử dụng nhánh **`main`** hoặc **`master`** làm nhánh ổn định, luôn sẵn sàng để deploy.
    - Tuyệt đối **không** commit trực tiếp lên nhánh chính.

2.  **Phát triển Tính năng/Sửa lỗi**:

    - Tạo nhánh mới từ nhánh **`main`** cho mỗi task/tính năng/sửa lỗi.
    - Quy ước đặt tên nhánh:
      - Tính năng mới: `feature/<tên-tính-năng-ngắn>`
      - Sửa lỗi: `fix/<tên-lỗi-ngắn>`
    - _Ví dụ_: `feature/user-profile-update`, `fix/login-bug-404`.

3.  **Yêu cầu Hợp nhất (Pull/Merge Request)**:

    - Khi hoàn thành, tạo **Pull Request (PR)** từ nhánh của bạn vào nhánh **`develop`**.
    - **Bắt buộc** phải có ít nhất **một** reviewer chấp thuận trước khi merge.
    - Sử dụng tùy chọn **Squash and Merge** khi merge để giữ lịch sử khi commit trên nhánh `main` và nhánh `develop` gọn gàng, mỗi PR tương ứng với một commit duy nhất tuân thủ quy chuẩn commit.

4.  **Kiểm tra và Tự động hóa (CI/CD)**:

    - Mọi PR đều **tự động** chạy **Continuous Integration (CI)** (test, lint, build check). Chỉ được phép merge khi CI **thành công**.

---

## Quy chuẩn Code & Best Practices

Để duy trì chất lượng và sự đồng nhất của code trên toàn bộ hệ thống, chúng ta áp dụng các quy tắc sau:

### 1\. Style Guide (Chủ yếu)

> Một số ngôn ngữ có chuẩn code lạ, tham khảo phần [ngoại lệ](#2-ngoại-lệ)

- **Đồng nhất:** Sử dụng **Linter** (ví dụ: ESLint, Rustfmt, Gofmt) và **Formatter** (ví dụ: Prettier) để tự động hóa việc định dạng code.
  - _Mọi dự án đều phải có file cấu hình Linter/Formatter._
- **Đặt tên:**
  - Sử dụng **camelCase** cho biến, hàm và phương thức (ví dụ: `calculateTotalPrice`).
  - Sử dụng **PascalCase** cho Class, Component và Type (ví dụ: `UserFactory`, `HeaderComponent`).
  - Sử dụng **SCREAMING_SNAKE_CASE** cho hằng số (ví dụ: `MAX_RETRIES`).

### 2\. Ngoại lệ

#### 1. Ngôn ngữ Rust (Rust Naming Conventions)

Rust có các quy tắc đặt tên riêng biệt và nghiêm ngặt để phân biệt rõ ràng giữa các loại cấu trúc (struct), enum, hàm, biến và hằng số.

| Cấu trúc (Construct)    | Quy tắc Đặt tên                               | Ví dụ                                              |
| :---------------------- | :-------------------------------------------- | :------------------------------------------------- |
| **Biến (Variables)**    | **`snake_case`** (Chữ thường, dùng gạch dưới) | `let max_attempts = 5;`                            |
| **Hàm (Functions)**     | **`snake_case`**                              | `fn calculate_total_price() -> f64 { ... }`        |
| **Tham số Hàm**         | **`snake_case`**                              | `fn process_data(user_input: &str) { ... }`        |
| **Mô-đun (Modules)**    | **`snake_case`**                              | `mod database_utils { ... }`                       |
| **Struct, Enum**        | **`PascalCase`** (Giống Class)                | `struct UserProfile { ... }`, `enum State { ... }` |
| **Trait, Type Alias**   | **`PascalCase`**                              | `trait Flyable { ... }`, `type Result<T> = ...`    |
| **Hằng số (Constants)** | **`SCREAMING_SNAKE_CASE`**                    | `const MAX_RETRIES: i32 = 10;`                     |

### 3\. Xử lý Lỗi (Error Handling)

- **Hạn chế `panic`/crash:** Trừ các lỗi không thể phục hồi (ví dụ: lỗi khởi tạo hệ thống), các lỗi nghiệp vụ cần được xử lý bằng cách trả về `Result` (Rust/Golang) hoặc sử dụng cơ chế `try...catch` (React/Flutter) rõ ràng.
- **Logging:** Sử dụng cấu trúc **[Structured Logging](https://www.loggly.com/use-cases/what-is-structured-logging-and-how-to-use-it/)** (JSON/Key-Value) để dễ dàng tìm kiếm và phân tích lỗi. Cấp độ log phải được xác định rõ ràng (DEBUG, INFO, WARN, ERROR).

### 3\. Testing (Kiểm thử)

- **Code Coverage:** Mọi tính năng mới và sửa lỗi lớn đều phải đi kèm với **Unit Tests** và **Integration Tests** tương ứng. Mục tiêu duy trì **code coverage tối thiểu là 80%**.
- **Mocking:** Sử dụng các thư viện Mocking/Stubbing để đảm bảo Unit Tests chỉ kiểm tra logic của một đơn vị code độc lập, loại bỏ phụ thuộc bên ngoài.

---

## Ví dụ / Tutorial

Thư mục `examples/` chứa các ví dụ nhỏ, tự chứa, giúp bạn nhanh chóng làm quen với **cấu trúc thư mục**, **cách đặt tên**, và **best practices** cho từng ngôn ngữ/framework:

- **examples/rust**: Ví dụ về cấu trúc một Microservice Rust điển hình.
- **examples/golang**: Ví dụ về cấu trúc Golang API với error handling.
- **examples/flutter**: Ví dụ về cách quản lý state và cấu trúc thư mục.
- **examples/react**: Ví dụ về tổ chức component và sử dụng hooks.
