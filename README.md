# Khởi đầu của Overlix

Chào mừng bạn đến với repository onboarding nội bộ

Repository này giúp **nhân viên mới nhanh chóng hiểu kiến trúc hệ thống, chuẩn code, workflow Git và best practices** của công ty.

> ⚠️ Tài liệu này chưa đầy đủ và đang được bổ sung.

---

## Mục Lục

- [Bắt đầu nhanh](#bắt-đầu-nhanh)
- [Tổng quan kiến trúc](#tổng-quan-kiến-trúc)
- [Quy chuẩn Git & Commit](#quy-chuẩn-git--commit)

---

## Bắt đầu nhanh

1. Clone repository:

```bash
git clone https://git.overlix.net/overlix/onboarding.git
cd onboarding
```

2. Tham khảo các ví dụ trong thư mục `examples/`:

- Rust: `examples/rust`
- Golang: `examples/golang`
- Flutter: `examples/flutter`
- React: `examples/react`

3. Đọc tài liệu module trong `ARCHITECTURE.md` để hiểu cấu trúc hệ thống.

---

## Sơ lược kiến trúc

Hệ thống bao gồm:

- **Application**: Ứng dụng dành cho người dùng cuối, chịu trách nhiệm cung cấp giao diện và trải nghiệm đầy đủ.
- **Website**: Trang web cho người dùng, phục vụ cho việc truy cập từ trình duyệt, cung cấp chức năng tương tự ứng dụng.
- **API**: Ứng dụng để giao diện người dùng tương tác với phía máy chủ.

---

## Quy chuẩn Git & Commit

Commit message cần theo cấu trúc như sau:

```text
<type>(optional scope): <description>

[optional body]

[optional footer]
```

- **fix:** Sửa lỗi (Tương đương với `PATCH` trong semantic versioning)
- **feat:** Thêm tính năng mới (Tương đương với `MINOR` trong semantic versioning)
- **refactor:** Refactor code, không thêm tính năng hay sửa lỗi
- **chore:** Thay đổi hỗ trợ, cấu hình, không ảnh hưởng chức năng
- **test:** Thêm hoặc sửa test
- **docs:** Các thay đổi chỉ liên quan đến documentation
- **perf:** Các thay đổi code nhằm tăng hiệu năng
- **ci:** thay đổi đối với configuration và scripts file CI
- **BREAKING CHANGE:** Một commit có chữ `BREAKING CHANGE:` ở phần đầu của `[optional body]`
  hoặc phần footer sẽ giới thiệu những thay đổi API bị phá vỡ (Tương đương với `MAJOR` trong semantic versioning). MỘT BREAKING CHANGE
  có thể là một phần commits của bất kì `<type>` nào

### Ví dụ

#### Commit message với mô tả và breaking change trong body

```
feat: thêm tính năng subgroup vào api

BREAKING CHANGE: tính năng `subgroup` giới thiệu nhiều endpoints mới
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

## Ví dụ / Tutorial

Thư mục `exmaples/` chứa các ví dụ nhỏ, tự chứa:
